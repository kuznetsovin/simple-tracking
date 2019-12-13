import axios from "axios";

let AddGeoObjectWidget = function () {
    let _addGeoObjectBtn = document.getElementById("addGeoObjectBtn");
    let _saveGeoObjectBtn = document.getElementById("saveGeoObjectBtn");
    let _closeModalBtn = document.getElementById("closeObjModalBtn");
    let _newGeoObjectModal = document.getElementById("addGeoObjectModal");

    function _openModal() {
        _newGeoObjectModal.classList.add("show");
        _newGeoObjectModal.style.display = "block";
    }

    function _closeModal() {
        _newGeoObjectModal.classList.remove("show");
        _newGeoObjectModal.style.display = "none";
    }

    this.init = function () {
        // привязываем обработчик показа модального окна
        _addGeoObjectBtn.addEventListener("click", () => {
            _openModal();
        });

        // привязываем обработчик сохранения ТС
        _saveGeoObjectBtn.addEventListener("click", () => {
            let geObjectName = document.getElementById("newGeoObjectName");
            let geObjectGeom = document.getElementById("newGeoObjectGeom");

            axios.post(BACKEND_URL + "/geo-objects",
                {name: geObjectName.value, geom: JSON.parse(geObjectGeom.value)}).then(
                function (response) {
                    if (response.status === 201) {
                        _closeModal();
                        geObjectName.value = '';
                        geObjectGeom.value = '';
                    }
                }).catch((err) => console.log(err));
        });

        // привязываем обработчик закрытия модального окна
        _closeModalBtn.addEventListener("click", () => {
            _closeModal();
        });
    }
};

export default AddGeoObjectWidget;