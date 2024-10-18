interface ProgramCardProps {
    name: string;
    split: string;
    perWeek: number;
}

const ProgramCard = (props: ProgramCardProps) => {
    return (
        <div>
            <h2> {props.name} </h2>
            <div>
                <p> split: {props.split} </p>
                <p> {props.perWeek} x Week </p>
            </div>
        </div>
    )
}

export default ProgramCard
