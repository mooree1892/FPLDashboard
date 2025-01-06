export interface Player {
    id: number;
    name: string;
    team: number;
    position: string;
    currentPrice: number;
    totalPoints: number;
    selectedBy: number;
    form: number;
    expectedPoints: number;
  }