export interface TransferSuggestion {
    playerId: number;
    playerName: string;
    currentTeam: number;
    transferScore: number;
    recommendedAction: 'STRONG_BUY' | 'CONSIDER' | 'HOLD';
  }