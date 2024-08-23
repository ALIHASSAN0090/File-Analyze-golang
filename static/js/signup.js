// scripts.js

function toggleForm(form) {
    const loginForm = document.getElementById('login-form');
    const signupForm = document.getElementById('signup-form');
    const loginToggle = document.getElementById('login-toggle');
    const signupToggle = document.getElementById('signup-toggle');

    if (form === 'login') {
        loginForm.classList.add('active');
        signupForm.classList.remove('active');
        loginToggle.classList.add('active');
        signupToggle.classList.remove('active');
    } else {
        signupForm.classList.add('active');
        loginForm.classList.remove('active');
        signupToggle.classList.add('active');
        loginToggle.classList.remove('active');
    }
}

// Set default form to login
document.addEventListener('DOMContentLoaded', () => {
    toggleForm('login');
});

async function handleLogin(event) {
    event.preventDefault();
    const username = document.getElementById('login-username').value;
    const password = document.getElementById('login-password').value;

    try {
        const response = await fetch('http://localhost:3000/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams({
                Name: username,
                Password: password,
            }),
        });

        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('token', data.token);
            window.location.href = '/home'; // Redirect to /home route
        } else {
            const errorData = await response.json();
            alert(`Login failed: ${errorData.Error}`);
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred');
    }
}

async function handleSignup(event) {
    event.preventDefault();
    const username = document.getElementById('signup-username').value;
    const password = document.getElementById('signup-password').value;

    try {
        const response = await fetch('http://localhost:3000/signup', { // Corrected URL
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams({
                Name: username,
                Password: password,
            }),
        });

        if (response.ok) {
            alert('Signup successful');
            toggleForm('login');
        } else {
            const errorData = await response.json();
            alert(`Signup failed: ${errorData.Error}`);
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred');
    }
}

// Function to include token in requests to protected routes
async function fetchWithToken(url, options = {}) {
    const token = localStorage.getItem('token');
    if (!token) {
        throw new Error('No token found');
    }

    options.headers = {
        ...options.headers,
        'Authorization': `Bearer ${token}`,
    };

    const response = await fetch(url, options);
    return response;
}


async function fetchProtectedData(url) {
    try {
        const response = await fetchWithToken(url);
        if (response.ok) {
            const data = await response.json();
            console.log('Protected data:', data);
        } else {
            console.error('Failed to fetch protected data');
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

// Call this function to fetch protected data after login
document.addEventListener('DOMContentLoaded', () => {
    if (localStorage.getItem('token')) {
        fetchProtectedData('http://localhost:3000/home');
        fetchProtectedData('http://localhost:3000/index.html');
        fetchProtectedData('http://localhost:3000/getAll.html');
        fetchProtectedData('http://localhost:3000/UpdateRecord.html');
        fetchProtectedData('http://localhost:3000/deleteRecord.html');
        fetchProtectedData('http://localhost:3000/createTable.html');
    }
});