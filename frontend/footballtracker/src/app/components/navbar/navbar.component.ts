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
    isLoggedIn: boolean;

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _router: Router
    ) { }

    ngOnInit() {
        this._footballService.currentMessage.subscribe(message => this.loginStatus = message);
        this._footballService.usernameMessage.subscribe(message => this.username = message);
        this.isLoggedIn = !this._footballService.isAuthenticated();

        if (this.isLoggedIn) {
            this._footballService.getCurrentUser().then(response => {
                sessionStorage.setItem('username', response.data.username);
                sessionStorage.setItem('userId', response.data.userId);
                    
                this._footballService.changeMessage('true');
                this.username = sessionStorage.getItem('username');
                this.userId = Number(sessionStorage.getItem('userId'));
            });
        } else {
            this._footballService.changeMessage('false');
        }
    }

    /**
     * Set new page title
     * @param {string} newTitle New page title
     */
    setProperties(newTitle: string) {
        this._titleService.setTitle(newTitle);
    }

    /**
     * Log user out of the platform
     */
    logout() {
        this._footballService.logout().then(response => {
            if (response.success) {
                sessionStorage.removeItem('username');
                sessionStorage.removeItem('userId');

                this.isLoggedIn = false;
                this._footballService.changeMessage('false');
                this._router.navigate(['/']);
            }
        })
    }
}