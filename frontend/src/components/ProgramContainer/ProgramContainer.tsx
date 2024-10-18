import { useState, useEffect } from 'react';
import Program from '../../types/program';
import ProgramCard from '../../components/ProgramCard/ProgramCard'

const ProgramContainer = () => {
  const [programs, setPrograms] = useState<Program[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchExercises = async () => {
      try {
        const response = await fetch('api/program', {
          method: 'GET',
          credentials: 'include', // Include cookies in the request
        });

        if (!response.ok) {
          throw new Error('Failed to fetch exercises');
        }

        const data = await response.json();
        setPrograms(data);
      } catch (error: any) {
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
    <div>
      <div>
        {programs.map((program: Program) => (
          <ProgramCard
            key={program.id}
            name={program.name}
            split={program.split}
            perWeek={program.per_week}
          />
        ))}
      </div>
    </div>
  );
};

export default ProgramContainer;


