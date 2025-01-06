import axios from 'axios';
import { Player } from '../types/player';
import { TransferSuggestion } from '../types/transfer';

class FPLApiService {
  private baseURL = 'http://localhost:8080/api';

  async getTransferSuggestions(): Promise<TransferSuggestion[]> {
    try {
      const response = await axios.get<TransferSuggestion[]>(`${this.baseURL}/transfers/suggestions`);
      return response.data;
    } catch (error) {
      console.error('Error fetching transfer suggestions:', error);
      throw error;
    }
  }

  async getPlayers(): Promise<Player[]> {
    try {
      const response = await axios.get<Player[]>(`${this.baseURL}/players`);
      return response.data;
    } catch (error) {
      console.error('Error fetching players:', error);
      throw error;
    }
  }

  // Additional methods for other API interactions can be added here
}

export default new FPLApiService();