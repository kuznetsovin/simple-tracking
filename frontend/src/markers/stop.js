import {Circle, Fill, Stroke, Style} from "ol/style";

let VehicleStopStyle = new Style({
    image: new Circle({
        radius: 7,
        fill: new Fill({color: 'grey'}),
        stroke: new Stroke({
            color: 'white', width: 2
        })
    })
});

export default VehicleStopStyle;