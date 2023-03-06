export const Login = (props) => {
    return (
        <div className="form">
            <button className='link-button' onClick={() => props.onFormSwitch('register')}>Create Account</button>
        </div>
    )
}