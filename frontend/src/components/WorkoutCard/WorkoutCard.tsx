interface WorkoutCardProps {
    name: string;
}

const ProgramCard = (props: WorkoutCardProps) => {
    return (
        <div>
            <h2> ${props.name} </h2>
        </div>
    )
}

export default ProgramCard

