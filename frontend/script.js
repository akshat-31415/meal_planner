const registerForm = document.getElementById('register-form');
const loginForm = document.getElementById('login-form');
const profileSection = document.getElementById('profile');

registerForm.addEventListener('submit', async (e) => {
  e.preventDefault();
  const name = document.getElementById('register-name').value;
  const email = document.getElementById('register-email').value;
  const password = document.getElementById('register-password').value;

  const res = await fetch('/users/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name, email, password })
  });

  const data = await res.json();
  alert(data.message);
});

loginForm.addEventListener('submit', async (e) => {
  e.preventDefault();
  const email = document.getElementById('login-email').value;
  const password = document.getElementById('login-password').value;

  const res = await fetch('/users/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });

  const data = await res.json();
  localStorage.setItem('token', data.token);
  fetchProfile();
});

async function fetchProfile() {
  const token = localStorage.getItem('token');
  if (!token) return;

  const payload = JSON.parse(atob(token.split('.')[1]));
  const userId = payload.sub || payload.id || payload.userId || payload.user_id;

  const res = await fetch(`/users/${userId}`, {
    headers: { 'Authorization': `Bearer ${token}` }
  });

  const data = await res.json();
  document.getElementById('profile-name').textContent = data.user.name;
  document.getElementById('profile-email').textContent = data.user.email;
  registerForm.style.display = 'none';
  loginForm.style.display = 'none';
  profileSection.style.display = 'block';
}

function logout() {
  localStorage.removeItem('token');
  profileSection.style.display = 'none';
  registerForm.style.display = 'block';
  loginForm.style.display = 'block';
}

if (localStorage.getItem('token')) fetchProfile();