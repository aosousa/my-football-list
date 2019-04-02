import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

// Components
import { IndexComponent } from '@components/index/index.component';
import { SignupComponent } from '@components/signup/signup.component';

// Services

const routes: Routes = [
    {
        path: 'signup',
        component: SignupComponent,
    },
    {
        path: '',
        component: IndexComponent
    }
]

@NgModule({
    imports: [
        RouterModule.forRoot(routes, {
            useHash: false,
            anchorScrolling: 'enabled'
        })
    ],
    exports: [
        RouterModule
    ]
})
export class AppRoutingModule{}