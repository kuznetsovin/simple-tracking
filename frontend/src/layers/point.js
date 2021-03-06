/*
File contains extended layers for display moving points.

Author: Igor Kuznetsov
Email: me@swe-notes.ru

(c) Copyright by Igor Kuznetsov.
*/

import VectorLayer from "ol/layer/Vector";
import {Vector} from "ol/source";
import {VehicleStopStyle} from "../markers/style";

let PointLayer = new VectorLayer({
    source: new Vector({})
});

PointLayer.addVehicleFromFeature = function (feature) {
    feature.setStyle(VehicleStopStyle);
    feature.setId(feature.get('client'));
    this.getSource().addFeature(feature);
};

PointLayer.getVehicleByID = function (clientId) {
    return this.getSource().getFeatureById(clientId);
};


export default PointLayer;