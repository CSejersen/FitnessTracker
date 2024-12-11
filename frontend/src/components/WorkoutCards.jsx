const Workouts = () => {
  const [workouts, setWorkouts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchExercises = async () => {
      try {
        const response = await fetch('api/workout', {
          method: 'GET',
          credentials: 'include', // Include cookies in the request
        });

        if (!response.ok) {
          throw new Error('Failed to fetch workouts');
        }

        const data = await response.json();
        setWorkouts(data);
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
      {workouts.map((workout) => (
        <div key={workout.id} className="bg-white shadow-md rounded-lg p-4">
          <h2 className="text-l font-bold">{workout.name}</h2>
          <p>{workout.id}</p>
        </div>
      ))}
      <div key={"new_workout"} className="bg-white shadow-md rounded-lg p-4">
        <h2 className="text-l font-bold">Add a new workout</h2>
      </div>
    </div>
  );
};

export default Workouts;
