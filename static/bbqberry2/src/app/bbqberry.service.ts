import {Injectable} from '@angular/core';

import {Observable} from 'rxjs/Observable';
import {of} from 'rxjs/observable/of';
import {catchError} from 'rxjs/operators';

import {HardwareConfig} from "./models/hardware";
import {MessageService} from './message.service';
import {HttpClient} from '@angular/common/http';
import {TemperatureReading} from "./models/temperature";
import {LightStrip} from "./models/lightstrip";

import {pipe} from 'rxjs/util/pipe';

@Injectable()
export class BbqberryService {
  private baseAPIURL = 'http://localhost:8080/api';
  private hwCfgURL = `${this.baseAPIURL}/hardware`;
  private grillLightsURL = `${this.baseAPIURL}/lights/grill`;
  private temperatureURL = `${this.baseAPIURL}/temperatures`;

  constructor(private http: HttpClient,
              private messageService: MessageService) {
  }

  private log(message: string) {
    this.messageService.add('BbqberryService: ' + message);
  }

  public getTemperatureReading(probe: number): Observable<TemperatureReading> {
    return this.http.get<HardwareConfig>(`${this.temperatureURL}?probe=${probe}`)
      .pipe(
        // tap(heroes => console.log(`getTemperatureReading`)),
        catchError(this.handleError('getTemperatureReading', null))
      );
  }

  public getHardwareConfig(): Observable<HardwareConfig> {
    return this.http.get<HardwareConfig>(this.hwCfgURL)
      .pipe(
        // tap(heroes => console.log(`getHardwareConfig`)),
        catchError(this.handleError('getHardwareConfig', null))
      );
  }

  public getGrillLightStrip(): Observable<LightStrip> {
    return this.http.get<LightStrip>(this.grillLightsURL)
      // Why the fuck does this not work!?!??!?!
      .retryWhen(attempts => Observable.range(1, 3)
        .zip(attempts, i => i)
        .mergeMap(i => {
          console.log("delay retry by " + i + " second(s)");
          return Observable.timer(i * 1000);
        })
      )
      .pipe(
        // tap(heroes => console.log(`getGrillLightStrip`)),
        catchError(this.handleError('getGrillLightStrip', null))
      );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      if (error != null) {
        console.error(error);
        this.log(`${operation} failed: ${error.message}`);
      }

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
