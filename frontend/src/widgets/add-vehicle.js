/*
File contains button widget for appending vehicle.

Author: Igor Kuznetsov
Email: me@swe-notes.ru

(c) Copyright by Igor Kuznetsov.
*/
import axios from "axios";

let AddVehicleWidget = function () {
    let _addVehicleBtn = document.getElementById("addVehicleBtn");
    let _saveVehicleBtn = document.getElementById("saveVehicleBtn");
    let _closeModalBtn = document.getElementById("closeModalBtn");
    let _newVehicleModal = document.getElementById("addVehicleModal");

    function _openModal() {
        _newVehicleModal.classList.add("show");
        _newVehicleModal.style.display = "block";
    }

    function _closeModal() {
        _newVehicleModal.classList.remove("show");
        _newVehicleModal.style.display = "none";
    }

    this.init = function () {
        // привязываем обработчик показа модального окна
        _addVehicleBtn.addEventListener("click", () => {
            _openModal();
        });

        // привязываем обработчик сохранения ТС
        _saveVehicleBtn.addEventListener("click", () => {
            let gosNumber = document.getElementById("newVehicleGosNumber");
            let gpsCode = document.getElementById("newVehicleGPS");

            axios.post(BACKEND_URL + "/vehicle-dict",
                {gos_number: gosNumber.value, gps_id: parseInt(gpsCode.value, 10)}).then(
                function (response) {
                    if (response.status === 201) {
                        _closeModal();
                        gosNumber.value = '';
                        gpsCode.value = '';
                    }
                }).catch((err) => console.log(err));
        });

        // привязываем обработчик закрытия модального окна
        _closeModalBtn.addEventListener("click", () => {
            _closeModal();
        });
    }
};

export default AddVehicleWidget;