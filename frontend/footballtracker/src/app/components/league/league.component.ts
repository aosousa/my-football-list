import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';
import { Response } from '@angular/http';

// 3rd party
import * as _ from 'lodash';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'league',
    templateUrl: './league.component.html',
    styleUrls: ['./league.component.css']
})
export class LeagueComponent implements OnInit {
    user: any = {}
    leagueFixtures: any = {};
    sessionUserId: number;
    leagueId: number;

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessagesService: FlashMessagesService,
        private _route: ActivatedRoute
    ) { }

    ngOnInit() {
        this.sessionUserId = Number(sessionStorage.getItem('userId'));
        this.leagueId = Number(this._route.snapshot.paramMap.get('id'));
        this.loadLeagueFixtures(this.leagueId);
        this._footballService.getUser(this.sessionUserId).then(response => {
            if (response.success) {
                this.user = response.data;
            }
        });
    }

    /**
     * Load a league's fixtures stored in the database
     * @param {number} leagueID ID of the league
     */
    loadLeagueFixtures(leagueID: number) {
        this._footballService.getLeagueFixtures(leagueID).then(response => {
            if (response.success && response.rows > 0) {
                this.leagueFixtures = response.data;
                this._titleService.setTitle("Football Tracker - " + this.leagueFixtures.league.name);
            } else {
                // league exists but has no fixtures
            }
        }).catch((error: Response) => {
            this._flashMessagesService.show(error.json().error, {
                cssClass: 'alert-danger',
                timeout: 10000
            });
        })
    }

        /**
     * Set fixture status as "watched" or "want to watch"
     * @param fixtureStatus 
     */
    setUserFixtureStatus(fixtureID, fixtureStatus, userFixtureId) {
        let userFixtureID = userFixtureId == 0 ? null : userFixtureId

        let userFixtureStatus = {
            fixtureId: fixtureID,
            status: fixtureStatus,
            userFixtureId: userFixtureID
        }

        this._footballService.createUserFixture(userFixtureStatus).then(response => {
            if (response.success) {
                this.loadLeagueFixtures(this.leagueId);
            }
        })
    }

    /**
     * Delete a user fixture status row
     * @param userFixtureId 
     */
    deleteUserFixture(userFixtureId) {
        this._footballService.deleteUserFixture(userFixtureId).then(response => {
            if (response.success) {
                this.loadLeagueFixtures(this.leagueId);
            }
        })
    }
}