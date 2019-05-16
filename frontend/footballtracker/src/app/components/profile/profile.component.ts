import { Component, OnInit, AfterViewInit, OnDestroy, ViewChildren, QueryList } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';
import { Subject } from 'rxjs';
import { DataTableDirective } from 'angular-datatables';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'profile',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit, AfterViewInit, OnDestroy {
    @ViewChildren(DataTableDirective)
    dtElements: QueryList<DataTableDirective>;

    user: any = {};
    userId: number;
    sessionUserId: number;
    userExists = true;
    dtOptions: DataTables.Settings = {};
    dtTrigger: Subject<any> = new Subject();
    public fixturesWatched: any = [];
    public fixturesInterestedIn: any = [];

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
        private _route: ActivatedRoute,
    ) { }

    ngOnInit() {
        this.dtOptions = {
            ordering: false,
            paging: true,
            pagingType: 'full_numbers',
            pageLength: 10,
            processing: true,
            lengthChange: false,
            destroy: true
        };

        this._titleService.setTitle("Football Tracker");
        this.userId = Number(this._route.snapshot.paramMap.get('id'));
        this.sessionUserId = Number(sessionStorage.getItem('userId'));
        this.loadUserFixtures(this.userId);

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

    ngAfterViewInit(): void {
        this.dtTrigger.next();
    }

    ngOnDestroy(): void {
        // Do not forget to unsubscribe the event
        this.dtTrigger.unsubscribe();
    }

    rerender(): void {
        this.dtElements.forEach((dtElement: DataTableDirective) => {
            dtElement.dtInstance.then((dtInstance: DataTables.Api) => {
                // Destroy the table first
                dtInstance.destroy();

                // Call the dtTrigger to re-render again
                this.dtTrigger.next();
            });
        });
    }

    /**
     * Load user's watched and interested in watching fixtures
     * @param {number} userID ID of the user
     */
    loadUserFixtures(userID: number) {
        this._footballService.getUserFixtures(userID).then(response => {
            if (response.success) {
                this.fixturesWatched = response.data.watched;
                this.fixturesInterestedIn = response.data.interestedIn;
                this.rerender();
            }
        }).catch(error => {
            this._flashMessageService.show("Error getting user's fixtures. Please try again later.", {
                cssClass: 'alert-danger',
                timeout: 10000000000000
            });
        })
    }

    /**
     * Set fixture status as "watched" or "want to watch"
     * @param {number} fixtureID ID of the fixture in the platform
     * @param {number} fixtureStatus Fixture status (1 = Watched, 2 = Want to Watch)
     * @param {number} userFixtureId ID of the user fixture row (0 means a new one will be created, otherwise it's an update to the row with the ID received)
     * @param {number} position Position of the fixture in either the fixturesWatched or fixturesInterestedIn array
     * @param {number} type Type of fixture status (1 = Watched, 2 = Want to Watch)
     */
    setUserFixtureStatus(fixtureID: number, fixtureStatus: number, userFixtureId: number, position: number, type: number) {
        let userFixtureID = userFixtureId == 0 ? null : userFixtureId

        let userFixtureStatus = {
            fixtureId: fixtureID,
            status: fixtureStatus,
            userFixtureId: userFixtureID
        }

        this._footballService.createUserFixture(userFixtureStatus).then(response => {
            if (response.success) {
                if (type == 1) {
                    // from watched tab
                    this.fixturesWatched[position].status = fixtureStatus;
                    this.fixturesWatched[position].userFixtureId = response.data.userFixtureId;
                } else {
                    // from interested in watching tab
                    this.fixturesInterestedIn[position].status = fixtureStatus;
                    this.fixturesInterestedIn[position].userFixtureId = response.data.userFixtureId;
                }
            }
        })
    }

    /**
     * Delete a user fixture status row
     * @param {number} userFixtureId ID of the user fixture row to delete
     * @param {number} position Position of the fixture in either the fixturesWatched or fixturesInterestedIn array
     * @param {number} type Type of fixture status (1 = Watched, 2 = Want to Watch)
     */
    deleteUserFixture(userFixtureId: number, position: number, type: number) {
        this._footballService.deleteUserFixture(userFixtureId).then(response => {
            if (response.success) {
                if (type == 1) {
                    // from watched tab
                    this.fixturesWatched[position].status = 0;
                    this.fixturesWatched[position].userFixtureId = 0;
                } else {
                    // from interested in watching tab
                    this.fixturesInterestedIn[position].status = 0;
                    this.fixturesInterestedIn[position].userFixtureId = 0;
                }
            }
        })
    }
}