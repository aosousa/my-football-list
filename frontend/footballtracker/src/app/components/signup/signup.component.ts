import { Component, OnInit } from '@angular/core'
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators, AbstractControl } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// Services
import { FootballService } from '@services/football.service';

// import custom validator to validate that password and confirm password fields match
import { MustMatch } from '@helpers/must-match.validator';

@Component({
    selector: 'signup',
    templateUrl: './signup.component.html',
    styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
    registerForm: FormGroup;
    submitted = false;

    constructor(
        private _formBuilder: FormBuilder,
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
        private _router: Router,
    ) { 
        this.registerForm = this._formBuilder.group({
            username: ['', [Validators.required, Validators.pattern('^[a-zA-Z0-9 \'\-]+$')], this.usernameExistenceValidator.bind(this)],
            email: ['', [Validators.required, Validators.email, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')], this.emailExistenceValidator.bind(this)],
            password: ['', [Validators.required, Validators.minLength(6)]],
            confirmPassword: ['', Validators.required]
        }, {
            validator: MustMatch('password', 'confirmPassword')
        });
    }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Sign Up");
    }

    // convenience getter for easy access to form fields
    get f() {
        return this.registerForm.controls;
    }

     usernameExistenceValidator(control: AbstractControl) {
        let userInfo = {
            username: control.value
        };

        return this._footballService.checkUsernameExistence(userInfo).then(response => {
            return response.rows === 1 ? { usernameExists: true } : null
        });
    }

    emailExistenceValidator(control: AbstractControl) {
        let userInfo = {
            email: control.value
        };

        return this._footballService.checkEmailExistence(userInfo).then(response => {
           return response.rows === 1 ? { emailExists: true } : null
        });
    }

    signup() {
        this.submitted = true;

        // stop here if form is invalid
        if (this.registerForm.invalid) {
            console.log(this.registerForm.get('username').errors);
            return
        }

        this._footballService.signup(this.registerForm.value).then(response => {
            this.submitted = false;
            if (response.success) {
                sessionStorage.setItem('username', this.registerForm.value.username);
                sessionStorage.setItem('userId', response.data.userId);

                this._footballService.changeMessage('true');
                this._footballService.changeUsernameSource(this.registerForm.value.username);
                this._router.navigate(['/fixtures']);
            }
        }).catch(error => {
            this._flashMessageService.show('An error occurred while signing up. Please try again later.', {
                cssClass: 'alert-danger',
                timeout: 5000
            })
        });
    }
}