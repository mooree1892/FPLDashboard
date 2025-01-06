import React, { useState, useEffect } from 'react';
import FPLApiService from '../services/apiService.ts';
import { Player } from '../types/player.ts';

const PlayersPage: React.FC = () => {
  const [players, setPlayers] = useState<Player[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPlayers = async () => {
      try {
        const fetchedPlayers = await FPLApiService.getPlayers();
        setPlayers(fetchedPlayers);
        setLoading(false);
      } catch (err) {
        setError('Failed to fetch players');
        setLoading(false);
      }
    };

    fetchPlayers();
  }, []);

  if (loading) return <div className="text-center py-4">Loading players...</div>;
  if (error) return <div className="text-red-500 text-center py-4">{error}</div>;

  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <div className="container mx-auto">
        <h1 className="text-3xl font-bold mb-6">Players</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {players.map((player) => (
            <div 
              key={player.id} 
              className="bg-white shadow-md rounded-lg p-4 hover:shadow-lg transition-shadow"
            >
              <h2 className="text-xl font-semibold">{player.name}</h2>
              <div className="mt-2">
                <p>Team: {player.team}</p>
                <p>Position: {player.position}</p>
                <p>Total Points: {player.totalPoints}</p>
                <p>Current Price: £{player.currentPrice}m</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default PlayersPage;