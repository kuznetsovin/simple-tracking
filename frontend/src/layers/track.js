/*
File contains extended layers for display historical track.

Author: Igor Kuznetsov
Email: me@swe-notes.ru

(c) Copyright by Igor Kuznetsov.
*/

import VectorLayer from "ol/layer/Vector";
import {Vector} from "ol/source";
import Feature from "ol/Feature";
import LineString from "ol/geom/LineString";
import TrackStyle from "../markers/track";
import axios from "axios";

let TrackLayer = new VectorLayer({
    visible: false,
    source: new Vector({})
});

TrackLayer.addClientTrack = function (trackPoints, client) {
    this.getSource().clear();

    let trackGeom = new Feature({geometry: new LineString(trackPoints)});

    trackGeom.setStyle(TrackStyle);
    trackGeom.setId(client);
    this.getSource().addFeature(trackGeom);
};

TrackLayer.getTrackByClient = function (client) {
    return this.getSource().getFeatureById(client)
};

TrackLayer.extendTrackByClient = function (client, coords) {
    let track = TrackLayer.getTrackByClient(client);

    if (track) {
        track.getGeometry().appendCoordinate(coords);
    }
};

TrackLayer.loadTrack = function(client, startDate, endDate) {
    axios.get(BACKEND_URL + "/tracks/" + client,
        {params: {start_date: startDate.utc().format(), end_date: endDate.utc().format()}}).then(
        function (response) {
            let track = [];
            // отрисовываем трек за период
            response.data.track.forEach((coord) => {
                track.push(coord)
            });

            TrackLayer.addClientTrack(track, client);
            TrackLayer.setVisible(true);
        }).catch((err) => console.log(err));
};

export default TrackLayer;