import { useState } from 'react';

const AddProgramForm = () => {
  const [error, setError] = useState<string | null>(null)
  const [name, setName] = useState<string>('')
  const [split, setSplit] = useState<string>('')
  const [perWeek, setPerWeek] = useState<number>(0)

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await fetch('api/program', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, split, perWeek }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        const errorMessage = errorData.error || 'Failed to create program';
        throw new Error(errorMessage);
      }
    } catch (err: any) {
      setError(err.message);
    }
  };

  return (
    <div>
      <h2> New Program </h2>
      < form onSubmit={handleSubmit} >
        <div>
          <label htmlFor="Name"> Program name: </label>
          < input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        < div>
          <label htmlFor="split"> Split: </label>
          < input
            type="text"
            id="split"
            value={split}
            onChange={(e) => setSplit(e.target.value)}
            required
          />
        </div>
        < div className="mb-4" >
          <label htmlFor="perWeek"> Workouts per week: </label>
          < input
            type="number"
            id="perWeek"
            value={perWeek}
            onChange={(e) => setPerWeek(Number(e.target.value))}
            required
          />
        </div>
        {error && <p> {error} </p>}
        <button type="submit">
          Create Program
        </button>
      </form>
    </div>
  )
}

export default AddProgramForm
