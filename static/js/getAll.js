document.getElementById('getall').addEventListener('click', function() {
    fetch('/display')
        .then(response => response.json())
        .then(data => {
            console.log('Data received:', data); // Debug line
            const tableBody = document.querySelector('#dataTable tbody'); // Updated ID here
            tableBody.innerHTML = ''; // Clear existing table rows
            data.forEach(row => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${row.id || 'N/A'}</td>
                    <td>${row.vowels || 'N/A'}</td>
                    <td>${row.capital || 'N/A'}</td>
                    <td>${row.small || 'N/A'}</td>
                    <td>${row.spaces || 'N/A'}</td>
                `;
                tableBody.appendChild(tr);
            });
        })
        .catch(error => console.error('Error fetching data:', error));
});
