import axios from 'axios';
import GeoJSON from "ol/format/GeoJSON";

let GeoObjectWidget = function (geoObjectLayer) {
    let _elementID = "visibleObject";
    let _objectCheckbox, _helperLabel, _objectList;

    function _hideListObject() {
        _helperLabel.getElementsByTagName("i")[0].classList.replace("fa-caret-down", "fa-caret-right");
        _objectList.classList.remove("show")
    }

    function _showListObject() {
        _helperLabel.getElementsByTagName("i")[0].classList.replace("fa-caret-right", "fa-caret-down");
        _objectList.classList.add("show");
    }

    function _cleanListObject() {
        let child = _objectList.lastElementChild;
        while (child) {
            _objectList.removeChild(child);
            child = _objectList.lastElementChild;
        }
    }

    function _mutedText() {
        _helperLabel.classList.replace("text-light", "text-muted");
    }

    function _lightText() {
        _helperLabel.classList.replace("text-muted", "text-light");
    }

    function _render() {
        let widgets = document.getElementById("widgets");
        let objectsWidget = document.createElement("div");
        objectsWidget.classList.add('row');
        objectsWidget.innerHTML = `<div class="alert helper geo-objects">
            <div class="form-check">
                <input class="form-check-input" type="checkbox" value="" id="${_elementID}">
                <label class="form-check-label" for="${_elementID}">
                    <a href="#" id="objectsLabel" class="text-decoration-none text-muted"><small>Объекты</small> <i
                                class="fa fa-caret-right"></i></a>
                </label>
                <div class="collapse" id="collapseObjectList"></div>
            </div>
        </div>`;
        widgets.appendChild(objectsWidget);
    }

    function _createObjectCheckbox(objID, objName) {
        let objectDiv = document.createElement('div');
        objectDiv.className = 'form-check';
        let checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.value = objID;
        checkbox.id = `object${objID}`;
        checkbox.checked = true;
        checkbox.addEventListener("click", (e) => {
            if (e.target.checked) {
                geoObjectLayer.showObject(e.target.value);
            } else {
                geoObjectLayer.hideObject(e.target.value);
            }
        });

        let label = document.createElement('label');
        label.classList.add("form-check-label", "text-light");
        label.setAttribute("for", `object${objID}`);
        label.innerHTML = `<small>${objName}</small>`;
        objectDiv.append(checkbox, label);
        return objectDiv
    }

    this.init = function () {
        _render();

        _objectCheckbox = document.getElementById(_elementID);
        _helperLabel = _objectCheckbox.parentElement.getElementsByTagName("a")[0];
        _objectList = _objectCheckbox.parentElement.getElementsByClassName("collapse")[0];

        // отображение геозон на карте
        _objectCheckbox.addEventListener("click", (e) => {
            let visibleObjGeoms = e.target.checked;
            if (visibleObjGeoms) {
                axios.get(BACKEND_URL + "/geo-objects").then(
                    function (response) {
                        (new GeoJSON()).readFeatures(response.data).forEach((feature) => {
                            // составляем список геообъектов
                            let objID = feature.get('id');
                            let objName = feature.get('name');
                            let objectDiv = _createObjectCheckbox(objID, objName);

                            _objectList.appendChild(objectDiv);

                            geoObjectLayer.addObject(feature);
                        });

                        _lightText();
                        _showListObject();
                    }
                ).catch((err) => console.log(err))
            } else {
                _mutedText();
                _hideListObject();
                _cleanListObject();
            }
            geoObjectLayer.setVisible(visibleObjGeoms);
        });

        // отображение списка геозон в окошке
        _helperLabel.addEventListener("click", (e) => {
            e.preventDefault();
            if (_objectList.classList.contains("show")) {
                _hideListObject();
            } else {
                _showListObject();
            }

        });
    }


};

export default GeoObjectWidget;