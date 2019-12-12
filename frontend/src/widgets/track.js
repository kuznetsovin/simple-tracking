let TrackWidget = function (trackLayer) {
    let _trackCheckbox, _helperLabel;
    let _elementID = 'visibleTrack';

    function _mutedText() {
        _helperLabel.classList.replace("text-light", "text-muted");
    }

    function _lightText() {
        _helperLabel.classList.replace("text-muted", "text-light");
    }

    this.init = function () {
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