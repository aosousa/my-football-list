import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'profile',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
    user: any = {};
    userId: number;
    sessionUserId: number;
    userExists = true;

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
        private _route: ActivatedRoute,
    ) { }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker");
        this.userId = Number(this._route.snapshot.paramMap.get('id'));
        this.sessionUserId = Number(sessionStorage.getItem('userId'));

        this._footballService.getUser(this.userId).then(response => {
            if (response.success) {
                this.user = response.data;
                this._titleService.setTitle("Football Tracker - " + this.user.username);
            } 
        }).catch(error => {
            this.userExists = false;
            this._flashMessageService.show('User does not exist.', {
                cssClass: 'alert-danger',
                timeout: 1000000
            });
        });
    }
}