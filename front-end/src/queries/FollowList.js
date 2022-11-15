import * as React from 'react';

import {
	List
} from '@mui/material';

import { useParams } from 'react-router-dom';
import { list } from './Profile';
import { getFollowers } from './actions';

export default function FollowList(props) {
	const username = useParams()["username"];

	const [followList, setFollowList] = React.useState([]);
	React.useEffect(() => {
		getFollowers(username).then(data => setFollowList(data));
	}, [username]);

	return (
		<List sx={{ width: '100%', maxWidth: 360, bgcolor: 'background.paper' }}>
			{followList.map(user => list(user))}
		</List>
	)
}