import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';
import { Response } from '@angular/http';

// 3rd party
import * as _ from 'lodash';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'team',
    templateUrl: './team.component.html',
    styleUrls: ['./team.component.css']
})
export class TeamComponent implements OnInit {
    user: any = {};
    teamFixtures: any = {};
    sessionUserId: number;
    teamId: number;
    dataTable: any;

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessagesService: FlashMessagesService,
        private _route: ActivatedRoute
    ) { }

    ngOnInit() {
        this.sessionUserId = Number(sessionStorage.getItem('userId'));
        this.teamId = Number(this._route.snapshot.paramMap.get('id'));
        this.loadTeamFixtures(this.teamId);
        this._footballService.getUser(this.sessionUserId).then(response => {
            if (response.success) {
                this.user = response.data;
            }
        });
    }

    /**
     * Load a team's fixtures stored in the database
     * @param {number} teamID ID of the team
     */
    loadTeamFixtures(teamID: number) {
        this._footballService.getTeamFixtures(teamID).then(response => {
            if (response.success && response.rows > 0) {
                this.teamFixtures = response.data;
                this._titleService.setTitle("Football Tracker - " + this.teamFixtures.team.name);
            } else {
                // team exists but has no fixtures
            }
        }).catch((error: Response) => {
            this._flashMessagesService.show(error.json().error, {
                cssClass: 'alert-danger',
                timeout: 10000
            });
        });
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
                this.loadTeamFixtures(this.teamId);
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
                this.loadTeamFixtures(this.teamId);
            }
        })
    }
}