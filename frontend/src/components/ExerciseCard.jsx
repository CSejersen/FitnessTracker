import React, { useState, useEffect } from 'react';

const Exercises = () => {
  const [exercises, setExercises] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchExercises = async () => {
      try {
        const response = await fetch('api/exercise', {
          method: 'GET',
          credentials: 'include', // Include cookies in the request
        });

        if (!response.ok) {
          throw new Error('Failed to fetch exercises');
        }

        const data = await response.json();
        setExercises(data);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchExercises();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div className="p-4 space-y-4">
      {exercises.map((exercise) => (
        <div key={exercise.id} className="bg-white shadow-md rounded-lg p-4">
          <h2 className="text-xl font-bold">{exercise.name}</h2>
          <p>{exercise.description}</p>
        </div>
      ))}
    </div>
  );
};

export default Exercises;
