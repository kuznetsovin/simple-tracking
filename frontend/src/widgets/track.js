let TrackWidget = function (trackLayer) {
    let _trackCheckbox, _helperLabel;
    let _elementID = 'visibleTrack';

    function _mutedText() {
        _helperLabel.classList.replace("text-light", "text-muted");
    }

    function _lightText() {
        _helperLabel.classList.replace("text-muted", "text-light");
    }

    function _render() {
        let widgets = document.getElementById("widgets");
        let trackWidget = document.createElement("div");
        trackWidget.classList.add('row');
        trackWidget.innerHTML = `<div class="alert helper layers">
            <div class="form-check">
                <input class="form-check-input" type="checkbox" value="" id="${_elementID}" checked>
                <label class="form-check-label text-light" for="${_elementID}" id="trackLabel">
                    <small>Трек</small>
                </label>
            </div>
        </div>`;
        widgets.appendChild(trackWidget);
    }

    this.init = function () {
        _render();

        _trackCheckbox = document.getElementById(_elementID);
        _helperLabel = _trackCheckbox.parentElement.getElementsByTagName("label")[0];

        // отображение трека на карте
        _trackCheckbox.addEventListener("click", (e) => {
            let isTrackVisible = e.target.checked;
            if (isTrackVisible) {
                _lightText()
            } else {
                _mutedText()
            }

            trackLayer.setVisible(isTrackVisible);
        });
    }
};

export default TrackWidget;