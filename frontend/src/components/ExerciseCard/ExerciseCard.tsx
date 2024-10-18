interface ExerciseCardProps {
    name: string;
}

const ProgramCard = (props: ExerciseCardProps) => {
    return (
        <div>
            <h2> ${props.name} </h2>
        </div>
    )
}

export default ProgramCard

