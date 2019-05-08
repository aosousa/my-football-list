import { Component, OnInit, AfterViewInit, OnDestroy, ViewChild } from '@angular/core';
import { ActivatedRoute,  Router } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';
import { Response } from '@angular/http';
import { Subject } from 'rxjs';
import { DataTableDirective } from 'angular-datatables';

// 3rd party
import * as _ from 'lodash';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'team',
    templateUrl: './team.component.html',
    styleUrls: ['./team.component.css']
})
export class TeamComponent implements OnInit, AfterViewInit, OnDestroy {
    @ViewChild(DataTableDirective)
    dtElement: DataTableDirective;
  
    user: any = {};
    public teamFixtures: any = [];
    sessionUserId: number;
    teamId: number;
    dtOptions: DataTables.Settings = {};
    dtTrigger: Subject<any> = new Subject();
    dataIsLoaded = false;

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessagesService: FlashMessagesService,
        private _route: ActivatedRoute,
        private _router: Router
    ) { }

    ngOnInit() {
        this.dtOptions = {
            ordering: false,
            paging: true,
            pagingType: 'full_numbers',
            pageLength: 10,
            processing: true,
            lengthChange: false,
        };

        this.sessionUserId = Number(sessionStorage.getItem('userId'));
        this.teamId = Number(this._route.snapshot.paramMap.get('id'));
        this.loadTeamFixtures(this.teamId);
        this._footballService.getUser(this.sessionUserId).then(response => {
            if (response.success) {
                this.user = response.data;
            }
        });
    }

    ngAfterViewInit(): void {
        this.dtTrigger.next();
      }

      ngOnDestroy(): void {
        // Do not forget to unsubscribe the event
        this.dtTrigger.unsubscribe();
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

                this.dtElement.dtInstance.then((dtInstance: DataTables.Api) => {
                    // Destroy the table first
                    dtInstance.destroy();

                    // Call the dtTrigger to rerender again
                    this.dtTrigger.next();
                });
            } else {
                this.teamFixtures = response.data;

                // team exists but has no fixtures
                this._flashMessagesService.show("No matches found", {
                    cssClass: 'alert-danger',
                    timeout: 10000000000000
                });
            }
        }).catch((error: Response) => {
            this._flashMessagesService.show(error.json().error, {
                cssClass: 'alert-danger',
                timeout: 10000000000000
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

    changeTeam(teamID: number) {
        this._router.navigate(['/team', teamID]);
    }
}