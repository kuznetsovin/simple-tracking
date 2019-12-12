import axios from "axios";
import moment from "moment";
import "../asserts/css/report-table.css";

let ReportWidget = function (trackLayer) {
    let _reportCheckbox, _helperLabel, _reportPanel;
    let _elementID = 'visibleReportPanel';
    let _reportPanelID = 'report-panel';
    let _reportTableID = 'report-table';
    let _vehicleSelectID = "vehicleID";
    let _startDateID = "startDate";
    let _endDateID = "endDate";
    let _btnCreateID = "createTrack";
    let _dateFormat = "YYYY-MM-DDTHH:mm:ss";

    this.init = function () {
        _render();
        _createPanel();
        _loadVehicleList();

        let endDate = moment();
        let startDate = moment([endDate.year(), endDate.month(), endDate.date()]);

        this.setFilterValue(startDate, endDate);

        _reportCheckbox = document.getElementById(_elementID);
        _reportPanel = document.getElementById(_reportPanelID);
        _helperLabel = _reportCheckbox.parentElement.getElementsByTagName("label")[0];

        // отображение трека на карте
        _reportCheckbox.addEventListener("click", (e) => {
            let isPanelVisible = e.target.checked;
            if (isPanelVisible) {
                _showPanel()
            } else {
                _hidePanel()
            }
        });

        document.getElementById(_btnCreateID).addEventListener("click", () => {
            let startDate = moment(document.getElementById(_startDateID).value);
            let endDate = moment(document.getElementById(_endDateID).value);
            let client = document.getElementById(_vehicleSelectID).value;

            trackLayer.loadTrack(client, startDate, endDate);
            _getReportRows(client, startDate, endDate);
        })
    };

    this.setFilterValue = function (startDate, endDate, client) {
        document.getElementById(_endDateID).value = endDate.format(_dateFormat);
        document.getElementById(_startDateID).value = startDate.format(_dateFormat);
        document.getElementById(_vehicleSelectID).value = client
    };

    this.openPanel = function () {
        _reportCheckbox.checked = true;
        _showPanel();
    };

    function _showPanel() {
        _helperLabel.classList.replace("text-muted", "text-light");
        _reportPanel.style.zIndex = 3;
    }

    function _hidePanel() {
        _helperLabel.classList.replace("text-light", "text-muted");
        _reportPanel.style.zIndex = -1;
    }

    function _render() {
        let widgets = document.getElementById("widgets");
        let reportWidget = document.createElement("div");
        reportWidget.classList.add('row');
        reportWidget.innerHTML = `<div class="alert helper layers">
            <div class="form-check">
                <input class="form-check-input" type="checkbox" value="" id="${_elementID}">
                <label class="form-check-label text-muted" for="${_elementID}" id="trackLabel">
                    <small>Отчет</small>
                </label>
            </div>
        </div>`;
        widgets.appendChild(reportWidget);
    }

    function _createPanel() {
        let panelWidget = document.getElementById(_reportPanelID);
        panelWidget.innerHTML = `<h5>Параметры отчета</h5>
    <div class="content">
        <div class="form-group row">
            <label for="${_vehicleSelectID}" class="col-xs-1 col-form-label-sm">ТС</label>
            <div class="col-md-5">
                <select class="form-control form-control-sm" id="${_vehicleSelectID}"></select>
            </div>
        </div>
        <div class="form-group row">
            <label for="${_startDateID}" class="col-xs-1 col-form-label-sm">С&nbsp</label>
            <div class="col-md-5">
                <input type="datetime-local" class="form-control form-control-sm" id="${_startDateID}" value="">
            </div>
            <label for="${_endDateID}" ="col-xs-1 col-form-label-sm">&nbspпо&nbsp</label>
            <div class="col-md-5">
                <input type="datetime-local" class="form-control form-control-sm" id="${_endDateID}" value="">
            </div>
        </div>
        <div class="form-group row">
            <a href="#" class="btn btn-primary" id="${_btnCreateID}">Построить трек</a>
        </div>
        <div class="row">
            <table class="table table-striped table-bordered table-sm table-responsive-sm" id="${_reportTableID}">
                <thead>
                    <tr>
                        <th>Объект</th>
                        <th>Время первого заеда</th>
                        <th>Время последнего выезда</th>
                        <th>Пробег (км)</th>
                    </tr>
                </thead>
                <tbody></tbody>
            </table>
        </div>
    </div>`;
    }

    function _loadVehicleList() {
        let vehicleSelect = document.getElementById(_vehicleSelectID);
        axios.get(BACKEND_URL + "/vehicle-dict").then((response) => {
                // заполняем выпадающий список
                response.data.forEach((rec) => {
                    let option = document.createElement("option");
                    option.value = rec["gps_id"];
                    option.text = rec["gos_number"];
                    vehicleSelect.appendChild(option)
                });

            }
        ).catch((err) => console.log(err))
    }

    function _getReportRows(client, startDate, endDate) {
        axios.get(BACKEND_URL + "/report/object-dist/" + client,
            {params: {start_date: startDate.utc().format(), end_date: endDate.utc().format()}}).then(
            function (response) {
                _appendRowsToReportTable(response.data.report)
            }).catch((err) => console.log(err));
    }

    function _appendRowsToReportTable(rows) {
        let tblRow, cell_obj_name, cell_first_ts, cell_last_ts, cell_mileage;

        let reportTable = document.getElementById(_reportTableID).getElementsByTagName('tbody')[0];
        // очищаем записи кроме заголовка в таблице
        reportTable.innerHTML = '';

        rows.forEach((row, i) => {
            tblRow = reportTable.insertRow(i);

            cell_obj_name = tblRow.insertCell(0);
            cell_obj_name.innerHTML = `<small>${row["name"]}</small>`;

            cell_first_ts = tblRow.insertCell(1);
            cell_first_ts.innerHTML = `<small>${moment(row["first_point_timestamp"]).format("DD.MM.YYYY HH:mm:ss")}</small>`;

            cell_last_ts = tblRow.insertCell(2);
            cell_last_ts.innerHTML = `<small>${moment(row["last_point_timestamp"]).format("DD.MM.YYYY HH:mm:ss")}</small>`;

            cell_mileage = tblRow.insertCell(3);
            cell_mileage.innerHTML = `<small>${row["mileage"]}</small>`;
        })
    }
};

export default ReportWidget;