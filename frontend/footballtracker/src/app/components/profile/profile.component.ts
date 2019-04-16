import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
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
    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
        private _route: ActivatedRoute,
        private _router: Router
    ) { }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker");
        const userId = Number(this._route.snapshot.paramMap.get('id'));

        this._footballService.getUser(userId).then(response => {
            if (response.success) {
                this.user = response.data;
                this._titleService.setTitle("Football Tracker - " + this.user.username);
            } 
        }).catch(error => {
            // TODO: redirect and show flash message
        });
    }
}