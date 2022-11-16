import * as React from 'react';
import {
  List,
  Divider,
  ListItem,
  ListItemText,
  Avatar,
  ListItemAvatar
} from '@mui/material';
import { Check, NotInterested } from '@mui/icons-material';
import { useParams } from 'react-router-dom';
import { getUser } from './actions';

const list = (user) => {
	const fields = ["username", "id", "name", "description", "location", "url", "verified"];
	return fields.map(
		(field, index) => {
			return (
				<>
					<ListItem alignItems="flex-start">
						{field === "username" ? 
							<ListItemAvatar>
								<Avatar
									alt={user[field]}
									src={user["profile_image_url"]}
								/>
							</ListItemAvatar> : null}
						<ListItemText
							primary={field.toUpperCase()}
							primaryTypographyProps={{color: "primary"}}
							secondary={
							<React.Fragment>
								{user[field]}
							</React.Fragment>
							}
						/>
						{field === "verified" ? 
							<ListItemAvatar>
								<Avatar>{user[field] ? <Check/> : <NotInterested/>}</Avatar> 
							</ListItemAvatar> : null}
					</ListItem>
					{index !== fields.length-1 ? <Divider variant={field === "username" ? "inset" : "fullWidth"} component="li" /> : null}
				</>
			);
		}
	)
}

export default function UserProfile(props) {
	const username = useParams()["username"];

	const [user, setUser] = React.useState({});
	React.useEffect(() => {
		getUser(username).then(data => setUser(data));
	}, [username]);

	return (
		<List sx={{ width: '100%', maxWidth: 500, bgcolor: 'background.paper' }}>
			{user && list(user)}
		</List>
	)
}