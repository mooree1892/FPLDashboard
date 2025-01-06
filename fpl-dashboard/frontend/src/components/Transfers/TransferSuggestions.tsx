import React, { useState, useEffect } from 'react';
import { TransferSuggestion } from '../../types/transfer.ts';
import FPLApiService from '../../services/apiService.ts'

const TransferSuggestions: React.FC = () => {
  const [suggestions, setSuggestions] = useState<TransferSuggestion[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchTransferSuggestions = async () => {
      try {
        const transferSuggestions = await FPLApiService.getTransferSuggestions();
        setSuggestions(transferSuggestions);
        setLoading(false);
      } catch (err) {
        setError('Failed to fetch transfer suggestions');
        setLoading(false);
      }
    };

    fetchTransferSuggestions();
  }, []);

  const getActionColor = (action: string) => {
    switch (action) {
      case 'STRONG_BUY':
        return 'bg-green-100 text-green-800';
      case 'CONSIDER':
        return 'bg-yellow-100 text-yellow-800';
      case 'HOLD':
        return 'bg-gray-100 text-gray-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  if (loading) return <div className="text-center py-4">Loading suggestions...</div>;
  if (error) return <div className="text-red-500 text-center py-4">{error}</div>;

  return (
    <div className="bg-white shadow-md rounded-lg p-6">
      <h2 className="text-2xl font-bold mb-4">Transfer Suggestions</h2>
      <table className="w-full table-auto">
        <thead>
          <tr className="bg-gray-100">
            <th className="px-4 py-2">Player</th>
            <th className="px-4 py-2">Team</th>
            <th className="px-4 py-2">Transfer Score</th>
            <th className="px-4 py-2">Recommendation</th>
          </tr>
        </thead>
        <tbody>
          {suggestions.map((suggestion) => (
            <tr key={suggestion.playerId} className="border-b hover:bg-gray-50">
              <td className="px-4 py-2">{suggestion.playerName}</td>
              <td className="px-4 py-2">{suggestion.currentTeam}</td>
              <td className="px-4 py-2">{suggestion.transferScore.toFixed(2)}</td>
              <td className="px-4 py-2">
                <span className={`px-2 py-1 rounded ${getActionColor(suggestion.recommendedAction)}`}>
                  {suggestion.recommendedAction}
                </span>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TransferSuggestions;