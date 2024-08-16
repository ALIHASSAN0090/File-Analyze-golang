document.getElementById('analyzeForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission

    const formData = new FormData(this);

    fetch('/stats', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
       
        const tbody = document.querySelector('#resultsTable tbody');
        tbody.innerHTML = ''; // Clear any existing rows
        
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${data.total_vowels}</td>
            <td>${data.total_capitals}</td>
            <td>${data.total_small}</td>
            <td>${data.total_spaces}</td>
            <td>${data.process_time}</td>
        `;
        tbody.appendChild(row);
    })
    .catch(error => {
        console.error('Error:', error);
    });
});