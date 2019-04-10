import { Injectable } from '@angular/core';

@Injectable()
export class UtilsService {
    /**
     * Build a date in YYYY-mm-dd format.
     * @param {number} year Year in YYYY format
     * @param {number} month Number of the month
     * @param {number} day Day of the month
     */
    buildDate(year: number, month: number, day: number): string {
        return `${year}-${month}-${day}`;
    }
}