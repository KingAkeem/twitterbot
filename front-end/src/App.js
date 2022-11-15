import * as React from 'react';
import {
  CssBaseline,
  Drawer,
  Box,
  List,
  Divider,
  ListItem,
  ListItemButton,
  ListItemText,
  Typography,
  IconButton,
  Toolbar,
  AppBar as MuiAppBar,
  InputBase,
  Select,
  FormControl,
  InputLabel,
  MenuItem
} from '@mui/material';

import {
  Menu,
  ChevronLeft,
  ChevronRight,
  Search as SearchIcon,
} from '@mui/icons-material';

import { styled, useTheme, alpha } from '@mui/material/styles';

import './App.css';

import {
  Routes,
  Route,
  useNavigate,
} from "react-router-dom";

import UserProfile from './queries/user/Profile';

const Search = styled('form')(({ theme }) => ({
  position: 'relative',
  borderRadius: theme.shape.borderRadius,
  backgroundColor: alpha(theme.palette.common.white, 0.15),
  '&:hover': {
    backgroundColor: alpha(theme.palette.common.white, 0.25),
  },
  marginRight: theme.spacing(2),
  marginLeft: 0,
  width: '100%',
  [theme.breakpoints.up('sm')]: {
    marginLeft: theme.spacing(3),
    width: 'auto',
  },
}));

const SearchIconWrapper = styled('div')(({ theme }) => ({
  padding: theme.spacing(0, 2),
  height: '100%',
  position: 'absolute',
  pointerEvents: 'none',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
  color: 'inherit',
  '& .MuiInputBase-input': {
    padding: theme.spacing(1, 1, 1, 0),
    // vertical padding + font size from searchIcon
    paddingLeft: `calc(1em + ${theme.spacing(4)})`,
    transition: theme.transitions.create('width'),
    width: '100%',
    [theme.breakpoints.up('md')]: {
      width: '20ch',
    },
  },
}));

const Main = styled('main', { shouldForwardProp: (prop) => prop !== 'open' })(
  ({ theme, open }) => ({
    flexGrow: 1,
    padding: theme.spacing(3),
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    marginLeft: `-${drawerWidth}px`,
    ...(open && {
      transition: theme.transitions.create('margin', {
        easing: theme.transitions.easing.easeOut,
        duration: theme.transitions.duration.enteringScreen,
      }),
      marginLeft: 0,
    }),
  }),
);

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  transition: theme.transitions.create(['margin', 'width'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: `${drawerWidth}px`,
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: 'flex-end',
}));

const drawerWidth = 240;
function MainDrawer(props) {
  const theme = useTheme();
  const navigate = useNavigate();

  const [open, setOpen] = React.useState(false);
  const [searchText, setSearchText] = React.useState("");
  const [dataType, setDataType] = React.useState("user");

  const toggleDrawer =
    (open) =>
    (event) => {
      if (event.type === 'keydown' && ((event).key === 'Tab' || (event).key === 'Shift')) {
        return;
      }

      setOpen(open);
    };

  const body = () => (
    <>
      <DrawerHeader>
          <IconButton onClick={toggleDrawer(false)}>
            {theme.direction === 'ltr' ? <ChevronLeft /> : <ChevronRight />}
          </IconButton>
        </DrawerHeader>
        <Divider />
        <List>
          <ListItem key={"Query"} disablePadding>
            <ListItemButton>
              <ListItemText primary={"Query"} />
            </ListItemButton>
          </ListItem>
          <Divider />
          <ListItem key={"Export"} disablePadding>
            <ListItemButton>
              <ListItemText primary={"Export"} />
            </ListItemButton>
          </ListItem>
        </List>
      </>
  );

  return (
    <Box sx={{display: 'flex'}}>
      <CssBaseline/>
      <AppBar position="fixed" open={open}>
      <Toolbar>
        <IconButton
          color="inherit"
          aria-label="open drawer"
          onClick={toggleDrawer(true)}
          edge="start"
          sx={{ mr: 2, ...(open && { display: 'none' }) }}
        >
          <Menu />
        </IconButton>
        <Typography variant="h6" noWrap component="div">
          Search Social App
        </Typography>
        <Search onSubmit={e => {
            navigate(`/${dataType}/${searchText}`);
            e.preventDefault();
          }}>
            <SearchIconWrapper>
              <SearchIcon />
            </SearchIconWrapper>
            <StyledInputBase
              placeholder="Search…"
              inputProps={{ 'aria-label': 'search' }}
              value={searchText}
              onChange={e => {
                setSearchText(e.target.value);
              }}
            />
          </Search>
          <FormControl variant="filled">
            <InputLabel id="data-type">Data</InputLabel>
            <Select id="data-type" value={dataType} label="Data Type" onChange={e => setDataType(e.target.value)}>
              <MenuItem value="user">Profile</MenuItem>
              <MenuItem value="tweets">Tweets</MenuItem>
              <MenuItem value="followers">Followers</MenuItem>
              <MenuItem value="following">Following</MenuItem>
            </Select>
          </FormControl>
      </Toolbar>
    </AppBar>
    <Drawer
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        '& .MuiDrawer-paper': {
          width: drawerWidth,
          boxSizing: 'border-box',
        },
      }}
      variant="persistent"
      anchor={'left'}
      open={open}
    >
      {body()}
    </Drawer>
    <Main open={open}>
      <DrawerHeader/>
      {props.children}
      <Routes>
        <Route path="/user/:username" element={<UserProfile />}/>
      </Routes>
    </Main>
    </Box>
  );
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <MainDrawer/>
      </header>
    </div>
  );
}

export default App;