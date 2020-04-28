import {Component, ElementRef, HostListener, OnInit, Renderer2} from '@angular/core';

import {Probe} from "../models/hardware";
import {TemperatureReading} from "../models/temperature";
import {BbqberryService} from '../bbqberry.service';
import {Observable} from "rxjs/Rx";
import {SVGTransformer} from "./svg";

@Component({
  selector: 'app-temperature-indicator',
  // templateUrl: './temperature-indicator.component.html',
  template: '',
  styleUrls: ['./temperature-indicator.component.css'],
})
export class TemperatureIndicatorComponent implements OnInit {
  private elSVG: SVGElement;
  private elBody: SVGGElement;
  private elBodyOutline: SVGGElement;
  private elFill: SVGRectElement;
  private dHigh: DragableHandle;
  private dLow: DragableHandle;

  private padding: number = 50;
  private ticks: number = 10;
  private probeNumber: number = 0;

  private probe: Probe;
  private reading: TemperatureReading;
  private pollingData: any;

  constructor(private bbqService: BbqberryService,
              private renderer: Renderer2,
              private el: ElementRef) {
  }

  ngOnInit() {
    // Create the root SVG element
    this.elSVG = this.renderer.createElement('svg', 'svg');
    this.renderer.setAttribute(this.elSVG, 'width', (100 + this.padding * 2).toString());
    this.renderer.appendChild(this.el.nativeElement, this.elSVG);

    this.bbqService.getHardwareConfig()
      .subscribe(hwCfg => {
        if (hwCfg != null) {
          this.probe = hwCfg.probes[this.probeNumber];
          this.setupScale(this.probe)
        }
      });

    this.pollingData = Observable.interval(1000).startWith(0)
      .switchMap(() => this.bbqService.getTemperatureReading(0))
      .subscribe(readings => {
        if (readings != null) {
          this.adjustReading(readings[0]);
        }
      });
  }

  public ngOnDestroy() {
    this.pollingData.unsubscribe();
  }

  setupScale(probe: Probe) {
    this.probe = probe;

    this.renderer.setAttribute(this.elSVG, 'height', (this.probe.limits.maxAbsCelsius + this.padding * 2).toString());

    this.elBody = this.renderer.createElement('g', 'svg');
    this.elBody.setAttribute('id', `${this.probe.label}_elBody`);
    this.elBody.setAttribute('width', (Number(this.elSVG.getAttribute('width')) - this.padding * 2).toString());
    this.elBody.setAttribute('height', (Number(this.elSVG.getAttribute('height')) - this.padding * 2).toString());
    SVGTransformer.addTranslation(this.elBody, this.padding, this.padding);
    this.renderer.appendChild(this.elSVG, this.elBody);

    this.elBodyOutline = this.renderer.createElement('rect', 'svg');
    this.elBodyOutline.setAttribute('id', `${this.probe.label}_elBodyOutline`);
    this.elBodyOutline.setAttribute('x', '0');
    this.elBodyOutline.setAttribute('y', '0');
    this.elBodyOutline.setAttribute('width', this.elBody.getAttribute('width').toString());
    this.elBodyOutline.setAttribute('height', this.elBody.getAttribute('height').toString());
    this.renderer.setStyle(this.elBodyOutline, 'fill', 'white');
    this.renderer.setStyle(this.elBodyOutline, 'stroke', 'black');
    this.renderer.setStyle(this.elBodyOutline, 'stroke-width', '1');
    this.renderer.appendChild(this.elBody, this.elBodyOutline);

    this.fillTo(0);

    let height = Number(this.elBody.getAttribute("height"));
    let interval = height / this.ticks;
    for (let i=0, y=height; i <= this.ticks; i++, y-= interval) {
      let r = this.createTick(0, y, (height - y).toString());
      this.renderer.appendChild(this.elBody, r);
    }

    let hHigh = this.createLimitHandle(this.probe.limits.maxWarnCelsius, "high");
    this.renderer.appendChild(this.elBody, hHigh);
    this.dHigh = new DragableHandle(this.renderer, hHigh, this.elBodyOutline);

    let hLow = this.createLimitHandle(this.probe.limits.minWarnCelsius, "low");
    this.renderer.appendChild(this.elBody, hLow);
    this.dLow = new DragableHandle(this.renderer, hLow, this.elBodyOutline);
  }

  createLimitHandle(limit: number, type: string) {
    let w = 30;
    let h = 30;

    const limitHandle = this.renderer.createElement('g', 'svg');
    limitHandle.setAttribute('id', `${this.probe.label}_limitHandle_${type}`);
    SVGTransformer.addTranslation(limitHandle, -w, Number(this.elBody.getAttribute('height')) - limit - h/2);
    limitHandle.setAttribute( 'width', w.toString());
    limitHandle.setAttribute( 'height', h.toString());
    this.renderer.appendChild(this.elBody, limitHandle);

    const handlePoly = this.renderer.createElement('polygon', 'svg');
    handlePoly.setAttribute('fill', 'blue');
    handlePoly.setAttribute('points', `0,0 0,${h} ${w},${h/2}`);
    this.renderer.appendChild(limitHandle, handlePoly);

    const t = this.renderer.createElement('text', 'svg');
    t.setAttribute('fill', 'black');
    SVGTransformer.addTranslation(t, -2, 25);
    SVGTransformer.applyRotation(t, 270);
    t.innerHTML = limit.toString();
    this.renderer.appendChild(limitHandle, t);

    return limitHandle;
  }

  adjustReading(reading: TemperatureReading) {
    this.reading = reading;
    this.fillTo(reading.celsius);
  }

  fillTo(height: number) {
    if ( this.elFill == null ) {
      this.elFill = this.renderer.createElement('rect', 'svg');
      this.elFill.setAttribute('id', `${this.probe.label}_elFill`);
      this.elFill.setAttribute('x', '0');
      this.elFill.setAttribute('width', this.elBody.getAttribute('width').toString());
      this.renderer.setStyle(this.elFill, 'fill', 'red');
      this.renderer.appendChild(this.elBody, this.elFill);
    }
    this.elFill.setAttribute('y', (Number(this.elBody.getAttribute('height')) - height).toString());
    this.elFill.setAttribute('height', height.toString());
  }

  createTick(x, y: number, label: string) : any {
    const g = this.renderer.createElement('g', 'svg');
    SVGTransformer.addTranslation(g, x, y);

    const t = this.renderer.createElement('text', 'svg');
    t.setAttribute('x', '0');
    t.setAttribute('y', '0');
    t.setAttribute('fill', 'green');
    t.innerHTML = label;
    this.renderer.appendChild(g, t);

    const r = this.renderer.createElement('rect', 'svg');
    r.setAttribute( 'x', '0');
    r.setAttribute( 'y', '0');
    r.setAttribute( 'width', '10');
    r.setAttribute( 'height', '1');
    this.renderer.setStyle(r, 'fill', 'white');
    this.renderer.setStyle(r, 'stroke', 'green');
    this.renderer.setStyle(r, 'stroke-width', '1');
    this.renderer.setStyle(r, 'opacity', '1');
    this.renderer.appendChild(g, r);

    return g;
  }

  @HostListener('document:mousemove', ['$event'])
  onMouseMove(ev: MouseEvent) {
    if ( this.dHigh != null && this.dHigh.WantsMouseEvent(ev) ) {
      // console.log(`dHigh wants onMouseMove`);
      this.dHigh.onMouseMove(ev);
    }
    else if ( this.dLow != null && this.dLow.WantsMouseEvent(ev) ) {
      // console.log(`dLow wants onMouseMove`);
      this.dLow.onMouseMove(ev);
    }
    // else {
    //   console.log(`nobody wants onMouseMove`);
    // }
  }

  @HostListener('document:mousedown', ['$event'])
  onMouseDown(ev: MouseEvent) {
    if ( this.dHigh != null && this.dHigh.WantsMouseEvent(ev) ) {
      // console.log(`dHigh wants onMouseDown`);
      this.dHigh.onMouseDown(ev);
    }
    else if ( this.dLow != null && this.dLow.WantsMouseEvent(ev) ) {
      // console.log(`dLow wants onMouseDown`);
      this.dLow.onMouseDown(ev);
    }
    // else {
    //   console.log(`nobody wants onMouseDown`);
    // }
  }

  @HostListener('document:mouseup', ['$event'])
  onMouseUp(ev: MouseEvent) {
    if ( this.dHigh != null && this.dHigh.WantsMouseEvent(ev) ) {
      // console.log(`dHigh wants onMouseUp`);
      this.dHigh.onMouseUp(ev);
    }
    else if ( this.dLow != null && this.dLow.WantsMouseEvent(ev) ) {
      // console.log(`dLow wants onMouseUp`);
      this.dLow.onMouseUp(ev);
    }
    // else {
    //   console.log(`nobody wants onMouseUp`);
    // }
  }
}

class DragableHandle {
  private body: SVGElement;
  private el: SVGElement;
  private renderer: Renderer2;

  private dragging: boolean;
  private highLimit: number;
  private lowLimit: number;

  constructor(renderer: Renderer2, el: SVGElement, body: SVGElement) {
    this.dragging = false;

    this.el = el;
    this.body = body;
    this.renderer = renderer;

    // N.B:
    // Assumes the limit is half of our own element's height because im currently
    // using a triangle which has a point exactly in the right hand center of the client rect
    let rEl = this.el.getBoundingClientRect();
    let rBody = this.body.getBoundingClientRect();
    this.highLimit = rBody.top + rEl.height / 2;
    this.lowLimit = rBody.bottom - rEl.height / 2;

    this.el.addEventListener('mousedown', event => {this.onMouseDown(event);});
    this.el.addEventListener('mousemove', event => {this.onMouseMove(event);});
    this.el.addEventListener('mouseup', event => {this.onMouseUp(event);});
  }

  WantsMouseEvent(ev: MouseEvent) {
    if ( this.dragging ) {
      return true
    }

    let r = this.el.getBoundingClientRect();
    let insideX = ev.clientX > r.left && ev.clientX < r.right;
    let insideY = ev.clientY > r.top && ev.clientY < r.bottom;
    let inside = insideX && insideY;

    if ( inside && ev.type == "mousedown") {
      return true;
    }
  }

  onMouseDown(ev: MouseEvent) {
    this.dragging = true;
    // console.log(`onMouseDown ${event}`);
  }

  onMouseMove(ev: MouseEvent) {
    if ( this.dragging ) {
      let r = this.body.getBoundingClientRect();

      let tx = SVGTransformer.extractTranslation(this.el);

      // Easy case: completely inside the client rect
      let fullyInsideY = ev.clientY + ev.movementY > this.highLimit && ev.clientY + ev.movementY < this.lowLimit;
      if ( fullyInsideY ) {
        let newPos = [ tx[0], tx[1] + ev.movementY ];
        console.log(`A newPos[0]=${newPos[0]} newPos[0]=${newPos[1]} y=${ev.clientY} lowLimit=${this.lowLimit} highLimit=${this.highLimit}`);
        SVGTransformer.removeTranslation(this.el);
        SVGTransformer.addTranslation(this.el, newPos[0], newPos[1]);
      }
      // else {
      //   let newPos = [ tx[0], tx[1] ];
      //   if ( ev.clientY < this.highLimit ) {
      //     newPos[1] = this.highLimit;
      //     console.log(`B newPos[0]=${newPos[0]} newPos[0]=${newPos[1]} y=${ev.clientY} lowLimit=${this.lowLimit} highLimit=${this.highLimit}`);
      //   }
      //   else if ( ev.clientY > this.lowLimit ) {
      //     newPos[1] = this.lowLimit;
      //     console.log(`C newPos[0]=${newPos[0]} newPos[0]=${newPos[1]} y=${ev.clientY} lowLimit=${this.lowLimit} highLimit=${this.highLimit}`);
      //   }
      //   SVGTransformer.removeTranslation(this.el);
      //   SVGTransformer.addTranslation(this.el, newPos[0], newPos[1]);
      // }
    }
  }

  onMouseUp(ev: MouseEvent) {
    this.dragging = false;
    // console.log(`onMouseUp ${event}`);
    // let r = ;
    // console.log(`pos ${r}`);
  }
}

