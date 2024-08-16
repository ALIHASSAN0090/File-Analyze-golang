document.getElementById('updateForm').addEventListener('submit', function(event) {
    event.preventDefault();
    
    const formData = new FormData(this);
    const data = {
        id: formData.get('id'),
        value: formData.get('value')
    };

    fetch('/update', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(result => {
        alert(result.message || result.error);
    })
    .catch(error => console.error('Error:', error));
});