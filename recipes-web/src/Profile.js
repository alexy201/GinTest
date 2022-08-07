import React from 'react';
import './Profile.css'

const Profile = () => {
    //const { user, logout } = useAuth0();
    return (
        <li class="nav-item dropdown">
            
            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                <div class="user">
                    <span>{"LOGGED IN"}</span>
                </div>
            </a>
            <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                <a class="dropdown-item" /*onClick={() => logout()}*/> Logout</a>
            </div>
        </li>
    )
}

export default Profile;