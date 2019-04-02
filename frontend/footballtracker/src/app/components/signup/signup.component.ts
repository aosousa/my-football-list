import { Component, OnInit } from '@angular/core'
import { Router } from '@angular/router';
import { Title } from '@angular/platform-browser';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'signup',
    templateUrl: './signup.component.html',
    styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
    signupInfo: any;
    signupError: string;
    saving: boolean;

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _router: Router
    ) { }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Sign Up")
    }

    signup() {
        this.saving = true;

        this._footballService.signup(this.signupInfo).then(response => {
            this.saving = false;
            if (response.success) {
                this._router.navigate(['/fixtures']);
            } else {
                // display error message
                this.signupError = "An error occurred. Please try again later.";
            }
        })
    }
}