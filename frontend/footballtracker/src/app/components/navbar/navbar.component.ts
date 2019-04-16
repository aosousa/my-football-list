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
    userId: number;

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _router: Router
    ) { }

    ngOnInit() {
        this._footballService.currentMessage.subscribe(message => this.loginStatus = message);
        this._footballService.usernameMessage.subscribe(message => this.username = message);
        const loginStatus = this._footballService.isAuthenticated();

        if (loginStatus >= 0) {
            this._footballService.changeMessage('true');
            this.username = sessionStorage.getItem('username');
            this.userId = Number(sessionStorage.getItem('userId'));
        } else {
            this._footballService.changeMessage('false');
            this._router.navigate(['/']);
        }
    }

    setProperties(newTitle: string) {
        this._titleService.setTitle(newTitle);
    }

    logout() {
        this._footballService.logout().then(response => {
            if (response.success === true) {
                sessionStorage.removeItem('username');
                sessionStorage.removeItem('userId');

                this._footballService.changeMessage('false');
                this._router.navigate(['/']);
            }
        })
    }
}