import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot } from '@angular/router';
import { Injectable }  from '@angular/core';
import { Title } from '@angular/platform-browser';

// Services
import { FootballService } from '@services/football.service';

@Injectable()
export class AuthGuard implements CanActivate {
    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _router: Router 
    ) { }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
        return true;
    }
}