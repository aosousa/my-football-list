import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Title } from '@angular/platform-browser';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'navbar',
    templateUrl: './navbar.component.html',
    styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
    loginStatus: string;
    username: string;

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _router: Router
    ) { }

    ngOnInit() {
        this._footballService.currentMessage.subscribe(message => this.loginStatus = message);
        this._footballService.usernameMessage.subscribe(message => this.username = message);
        this.username = localStorage.getItem('username');
    }

    setProperties(newTitle: string) {
        this._titleService.setTitle(newTitle);
    }

    logout() {
        this._footballService.logout().then(response => {
            if (response.success === true) {
                localStorage.removeItem('username');

                this._router.navigate(['']);
                this._footballService.changeMessage('false');
            }
        })
    }
}