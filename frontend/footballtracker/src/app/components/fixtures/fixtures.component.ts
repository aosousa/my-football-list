import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// 3rd party
import * as _ from 'lodash';

// Services
import { FootballService } from '@services/football.service';
import { UtilsService } from '@services/utils.service';

@Component({
    selector: 'fixtures',
    templateUrl: './fixtures.component.html',
    styleUrls: ['./fixtures.component.css']
})
export class FixturesComponent implements OnInit {
    fixtures: any;
    groupedFixtures: any;
    dateVar: string;
    lastFixtureUpdate: string

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _utilsService: UtilsService,
        private _flashMessageService: FlashMessagesService
    ) { }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Fixtures");

        const now = new Date();
        this.dateVar = this._utilsService.buildDate(now.getFullYear(), _.padStart(String(now.getMonth() + 1), 2, '0'), _.padStart(String(now.getDate()), 2, '0'));

        this.loadFixtures(this.dateVar);
        this.getLastFixtureUpdate();
    }

    /**
     * Loads fixtures from a given day
     * @param {string} date Date in YYYY-mm-dd format
     */
    loadFixtures(date: string) {
        this._footballService.getFixturesByDate(date).then(response => {
            if (response.success) {
                this.fixtures = response.data;
            }
        }).catch(error => {
            this._flashMessageService.show('An error occurred while updating the fixtures list.', {
                cssClass: 'alert-danger',
                timeout: '5000'
            });
        });
    }

    /**
     * Get date/time of last fixture update
     */
    getLastFixtureUpdate() {
        this._footballService.getLastFixtureUpdate().then(response => {
            if (response.success) {
                this.lastFixtureUpdate = response.data;
            }
        })
    }
}