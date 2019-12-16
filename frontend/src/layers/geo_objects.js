/*
File contains extended layers for display geo object.

Author: Igor Kuznetsov
Email: me@swe-notes.ru

(c) Copyright by Igor Kuznetsov.
*/
import VectorLayer from "ol/layer/Vector";
import {Vector} from "ol/source";
import {EmptyStyle} from "../markers/style";

let GeoObjectLayer = new VectorLayer({
    visible: false,
    source: new Vector({})
});

GeoObjectLayer.addObject = function (feature) {
    let objectID = feature.get('id');
    feature.setId(objectID);
    this.getSource().addFeature(feature);
};

GeoObjectLayer.getObjectByID = function (objectID) {
    return this.getSource().getFeatureById(objectID);
};

GeoObjectLayer.hideObject = function (objectID) {
    let geom = this.getObjectByID(objectID);
    geom.setStyle(EmptyStyle);
};

GeoObjectLayer.showObject = function (objectID) {
    let geom = this.getObjectByID(objectID);
    geom.setStyle(null);
};


export default GeoObjectLayer;