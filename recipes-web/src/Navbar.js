import React from 'react';
import './Navbar.css'
import Profile from './Profile';

class Navbar extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            signedIn: false,
            recipes: []
        }

        this.checkAuthStatus();
        console.log(this.isAuthenticated);
    }

    checkAuthStatus() {
        var api_cookie = this.getCookie("recipes-api");
        this.signedIn = !(api_cookie == null);
    }

    getCookie(name) {
        var match = document.cookie.match(RegExp('(?:^|;\\s*)' + name + '=([^;]*)'));
        return match ? match[1] : null;
    }

    render() {
        return (
            <nav class="navbar navbar-expand-lg navbar-light bg-light">
                <a class="navbar-brand mb-0 h1" href="#">Recipes</a>
                <button class="navbar-toggler"
                    type="button"
                    data-toggle="collapse"
                    data-target="#navbarTogglerDemo02"
                    aria-controls="navbarTogglerDemo02"
                    aria-expanded="false"
                    aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse justify-content-end" id="navbarTogglerDemo02">
                    <ul class="navbar-nav ml-auto">
                        <li class="nav-item">
                            {this.isAuthenticated ? (<Profile />) : (
                                <a class="nav-link active" /*onClick={() => loginWithRedirect()}*/>
                                    Login</a>
                            )}
                        </li>
                    </ul>
                </div>
            </nav >
        )
    }
}


export default Navbar;

/*
const Navbar = () => {
    const { isAuthenticated, loginWithRedirect, logout, user } = useAuth0();
    return (
        <nav class="navbar navbar-expand-lg navbar-light bg-light">
            <a class="navbar-brand" href="#">Recipes</a>
            <button class="navbar-toggler"
                type="button" 
                data-toggle="collapse" 
                data-target="#navbarTogglerDemo02" 
                aria-controls="navbarTogglerDemo02" 
                aria-expanded="false" 
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse"
                id="navbarTogglerDemo02">
                <ul class="navbar-nav ml-auto">
                    <li class="nav-item">
                        {isAuthenticated ? (<Profile/>) : (
                            <a class="nav-link active"
                                onClick={() =>
                                    loginWithRedirect()}>
                                Login</a>
                        )}
                    </li>
                </ul>
            </div>
        </nav >
    )
}
export default Navbar;

*/