import 'ol/ol.css';
import '@fortawesome/fontawesome-free/css/all.min.css';
import './asserts/css/custom.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import {XYZ} from "ol/source";
import GeoJSON from "ol/format/GeoJSON";
import {VehicleMovingStyle} from "./markers/style";
import {GeoObjectLayer, PointLayer, TrackLayer} from "./layers/layers";
import {Map, View} from "ol";
import {Tile} from "ol/layer";

import axios from 'axios';
import moment from 'moment';
import TrackWidget from "./widgets/track";
import GeoObjectWidget from "./widgets/geo-object";
import ReportWidget from "./widgets/report";

let ws = new WebSocket(WEBSOCKET_URL);

let map = new Map({
    target: 'map',
    view: new View({
        projection: MAP_PROJ,
        center: MAP_CENTER,
        zoom: MAP_ZOOM,
        minZoom: 2,
        maxZoom: 19
    }),
    layers: [
        new Tile({
            source: new XYZ({url: MAP_SERVER_URL})
        }),
        PointLayer,
        TrackLayer,
        GeoObjectLayer,
    ]
});


map.once('postrender', function (event) {
    axios.get(BACKEND_URL + "/last-positions").then(
        function (response) {
            (new GeoJSON()).readFeatures(response.data).forEach((f) => {
                PointLayer.addVehicleFromFeature(f);
            });
        }
    );

    ws.onopen = function () {
        console.log("Connected");
    };

    ws.onmessage = function (event) {
        let currentFeature = (new GeoJSON()).readFeature(event.data);
        let client = currentFeature.get('client');
        let actualCoords = currentFeature.getGeometry().getCoordinates();

        let prevFeature = PointLayer.getVehicleByID(client);
        if (!prevFeature) {
            // странная ситуация, когда появляется ТС которой нет в последних точках
            console.log("Not found point: ", client);
            PointLayer.addVehicleFromFeature(currentFeature);
        }
        prevFeature.getGeometry().setCoordinates(actualCoords);
        prevFeature.setStyle(VehicleMovingStyle);

        TrackLayer.extendTrackByClient(client, actualCoords);

        map.render();
    };
});

window.onbeforeunload = function () {
    ws.close()
};

window.onload = function () {
    (new TrackWidget(TrackLayer)).init();
    (new GeoObjectWidget(GeoObjectLayer)).init();

    let reportPanel = new ReportWidget(TrackLayer);
    reportPanel.init();

    map.on("click", function (e) {
        map.forEachFeatureAtPixel(e.pixel, function (feature, layer) {

            // обработка клика по машинке
            map.getView().setCenter(feature.getGeometry().getCoordinates());
            map.getView().setZoom(18);

            let end_date = moment();
            let start_date = moment([end_date.year(), end_date.month(), end_date.date()]);
            let client = feature.id_;

            reportPanel.setFilterValue(start_date, end_date, client);
            reportPanel.openPanel();

            TrackLayer.loadTrack(client, start_date, end_date);
        });
    });
};