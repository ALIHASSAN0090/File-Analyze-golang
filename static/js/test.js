document.getElementById('testButton').addEventListener('click', function() {
    fetch('/test')
        .then(response => response.json())
        .then(data => {
            document.getElementById('responseStatus').textContent = data.working;
        })
        .catch(error => console.error('Error fetching status:', error));
});