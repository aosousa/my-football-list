import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators, AbstractControl } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'edit-profile',
    templateUrl: './edit-profile.component.html',
    styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit {
    user: any = {};
    userId: number;
    sessionUserId: number;
    editProfileForm: FormGroup;
    submitted = false;

    constructor(
        private _formBuilder: FormBuilder,
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
        private _route: ActivatedRoute,
        private _router: Router,
    ) { 
        this.editProfileForm = this._formBuilder.group({
            email: ['', [Validators.required, Validators.email, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')], this.emailExistenceValidator.bind(this)],
            spoilerMode: ['']        
        })
    }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Edit Profile")
        this.userId = Number(this._route.snapshot.paramMap.get('id'));
        this.sessionUserId = Number(sessionStorage.getItem('userId'));

        if (this.userId != this.sessionUserId) {
            this._flashMessageService.show('You are not authorized to perform this action.', {
                cssClass: 'alert-danger',
                timeout: 1000000
            });
        } else {
            this._footballService.getUser(this.userId).then(response => {
                if (response.success) {
                    this.user = response.data;
                    this.editProfileForm.patchValue({
                        email: this.user.email,
                        spoilerMode: this.user.spoilerMode
                    });
                }
            });
        }
    }

    // convenience getter for easy access to form fields
    get f() {
        return this.editProfileForm.controls;
    }

    emailExistenceValidator(control: AbstractControl) {
        let userInfo = {
            email: control.value
        };

        return this._footballService.checkEmailExistence(userInfo).then(response => {
            if (this.user.email == control.value) {
                return null
            } else {
                return response.rows === 1 ? { emailExists: true } : null
            }
        });
    }
    
    edit() {
        this.submitted = true;

        // stop here if form is invalid
        if (this.editProfileForm.invalid) {
            return
        }

        this._footballService.updateUser(this.userId, this.editProfileForm.value).then(response => {
            this.submitted = false;
            if (response.success) {
                this._flashMessageService.show('Your profile was updated successfully. Redirecting you to your profile page in 5 seconds.', {
                    cssClass: 'alert-success',
                    timeout: 5000
                });
                setTimeout(() => {
                    this._router.navigate(['/user/' + this.userId])
                }, 5000);
            } else {
                this._flashMessageService.show('An error occurred while trying to update your profile. Please try again later.', {
                    cssClass: 'alert-danger',
                    timeout: 10000
                });
            }
        }).catch(error => {
            this._flashMessageService.show('An error occurred while trying to update your profile. Please try again later.', {
                cssClass: 'alert-danger',
                timeout: 10000
            });
        })
    }
}