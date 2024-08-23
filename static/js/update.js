document.getElementById('updateForm').addEventListener('submit', function(event) {
    event.preventDefault();
    
    const formData = new FormData(this);
    const data = {
        id: formData.get('id'),
        value: formData.get('value')
    };

    console.log('Form Data:', data); // Add this line to log the data

    fetch('/update', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        console.log('Response Status:', response.status); // Log the status code
        return response.json();
    })
    .then(result => {
        console.log('Response Data:', result); // Log the response data
        alert(result.message || result.error);
    })
    .catch(error => {
        console.error('Fetch Error:', error); // Log fetch errors
    });
});
