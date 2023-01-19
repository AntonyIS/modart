import React, { Component } from 'react'

import axios from 'axios'

class Items extends Component {
    state = {
        users : [],
    };

    componentDidMount() {
        this.fetchusers()
    }

    async fetchusers(){
        const users = await axios.get("/api/v1/users");
        this.setState({users : users.data})
    }

    renderUsers(){
        return this.state.users.map(({firstname})=> firstname)
    }
    render (){
        <div>
            <h3>Users</h3>
            {
                this.renderUsers()
            }
        </div>
    }
}

export default Items;