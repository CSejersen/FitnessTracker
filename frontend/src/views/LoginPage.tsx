import { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom'; // Import Link here

const LoginPage = () => {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<string>('');
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await fetch('api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        const errorMessage = errorData.error || 'Login failed';
        throw new Error(errorMessage);
      }

      // Parse the JSON response and save token to local storage
      const data = await response.json();
      const token = data.token;
      localStorage.setItem('jwtToken', token);
      console.log('Login successful');
      navigate('/exercises');
    } catch (err: any) {
      setError(err.message);
    }
  };

  return (
    <div>
      <div>
        <h2> Login </h2>
        < form onSubmit={handleSubmit} >
          <div>
            <label htmlFor="username"> Username: </label>
            < input
              type="text" // Changed from 'username' to 'text'
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          < div className="mb-4" >
            <label htmlFor="password"> Password: </label>
            < input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          {error && <p> {error} </p>}
          <button type="submit">
            Login
          </button>
        </form>
        <div>
          <p>
            Don't have an account?{' '}
            < Link to="/signup">
              Sign up here
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
