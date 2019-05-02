import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

// Components
import { AboutComponent } from '@components/about/about.component';
import { ChangePasswordComponent } from '@components/change-password/change-password.component';
import { ContactComponent } from '@components/contact/contact.component';
import { EditProfileComponent } from '@components/edit-profile/edit-profile.component';
import { FixturesComponent } from '@components/fixtures/fixtures.component';
import { IndexComponent } from '@components/index/index.component';
import { LoginComponent } from '@components/login/login.component';
import { PrivacyPolicyComponent } from '@components/privacy-policy/privacy-policy.component';
import { ProfileComponent } from '@components/profile/profile.component';
import { ResetPasswordComponent } from '@components/reset-password/step1/reset-password.component';
import { NewPasswordComponent } from '@components/reset-password/step2/new-password.component';
import { SignupComponent } from '@components/signup/signup.component';
import { TeamComponent } from '@components/team/team.component';
import { TermsOfServiceComponent } from '@components/terms-of-service/terms-of-service.component';

// Services
import { AuthGuard } from '@services/auth.guard';

const routes: Routes = [
    {
        path: '',
        component: IndexComponent,
    },
    {
        path: 'signup',
        component: SignupComponent,
    },
    {
        path: 'login',
        component: LoginComponent
    },
    {
        path: 'about',
        component: AboutComponent
    },
    {
        path: 'privacy-policy',
        component: PrivacyPolicyComponent
    },
    {
        path: 'terms-of-service',
        component: TermsOfServiceComponent
    },
    {
        path: 'reset-password',
        component: ResetPasswordComponent
    },
    {
        path: 'password/:token',
        component: NewPasswordComponent
    },
    {
        path: 'fixtures',
        component: FixturesComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'contact',
        component: ContactComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'user/:id',
        component: ProfileComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'user/:id/edit',
        component: EditProfileComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'user/:id/change-password',
        component: ChangePasswordComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'team/:id',
        component: TeamComponent,
        canActivate: [AuthGuard]
    }
]

@NgModule({
    imports: [
        RouterModule.forRoot(routes, {
            useHash: false,
            anchorScrolling: 'enabled',
            onSameUrlNavigation: 'reload'
        })
    ],
    exports: [
        RouterModule
    ]
})
export class AppRoutingModule{}