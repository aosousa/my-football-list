import { Component, OnInit } from '@angular/core'
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators, ValidatorFn, AbstractControl, AbstractControlOptions, AsyncValidatorFn } from '@angular/forms';
import { Title } from '@angular/platform-browser';

import { of } from 'rxjs';
import { map } from 'rxjs/operators';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'signup',
    templateUrl: './signup.component.html',
    styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
    registerForm: FormGroup;
    submitted = false;
    debouncer: any;

    constructor(
        private _formBuilder: FormBuilder,
        private _titleService: Title,
        private _footballService: FootballService,
        private _router: Router,
    ) { 
        this.registerForm = this._formBuilder.group({
            username: ['', [Validators.required], this.usernameExistenceValidator.bind(this)],
            email: ['', [Validators.required, Validators.email, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')], this.emailExistenceValidator.bind(this)],
            password: ['', [Validators.required, Validators.minLength(6)]]
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
            return
        }

        this._footballService.signup(this.registerForm.value).then(response => {
            this.submitted = false;
            if (response.success) {
                // TEMPORARY: redirect to /fixtures when that route exists and perform login
                this._router.navigate(['/']);
            } 
        });
    }
}