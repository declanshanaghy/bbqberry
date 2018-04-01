import { TestBed, inject } from '@angular/core/testing';

import { BbqberryService } from './bbqberry.service';

describe('BbqberryService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [BbqberryService]
    });
  });

  it('should be created', inject([BbqberryService], (service: BbqberryService) => {
    expect(service).toBeTruthy();
  }));
});
