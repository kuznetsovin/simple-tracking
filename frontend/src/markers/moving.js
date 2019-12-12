import {Circle, Fill, Stroke, Style} from "ol/style";

let VehicleMovingStyle = new Style({
    image: new Circle({
        radius: 7,
        fill: new Fill({color: 'green'}),
        stroke: new Stroke({
            color: 'white', width: 2
        })
    })
});

export default VehicleMovingStyle;