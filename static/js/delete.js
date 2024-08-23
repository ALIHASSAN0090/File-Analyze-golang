document.querySelector('#deleteForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent default form submission

    const id = document.querySelector('input[name="id"]').value;
    console.log('Deleting record with ID:', id);

    fetch(`/delete?id=${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    })
    .then(response => response.json())
    .then(data => {
        console.log('Response data:', data);
        if (data.error) {
            alert('Error: ' + data.error);
        } else {
            alert('Success: ' + data.message);
            document.querySelector('input[name="id"]').value = '';
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
});
